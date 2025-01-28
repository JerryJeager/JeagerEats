"use client";

import { Minus, Plus, ShoppingCart, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { FaShoppingCart } from "react-icons/fa";
import useCartStore from "@/store/useCartStore";
import { PlaceOrderType } from "@/types";
import { getCookie } from "@/actions/handleCookies";
import { useState } from "react";
import { useToast } from "@/hooks/use-toast";
import Link from "next/link";
import { ToastAction } from "@/components/ui/toast";
import axios from "axios";
import { BASE_URL } from "@/data";
import { useRouter } from "next/navigation";
import Spinner from "../ui/Spinner";
import PaymentButton from "./PaymentButton";

export default function CartDialog() {
  const {
    menu,
    incrementQuantity,
    decrementQuantity,
    removeFromMenu,
    resetCart,
  } = useCartStore((state) => state);
  const router = useRouter();

  const total = menu.reduce((sum, item) => sum + item.price * item.quantity, 0);
  const deliveryFee = total * 0.05;
  const [isLoading, setIsLoading] = useState(false);
  const [isOpen, setIsOpen] = useState(false);
  const [email, setEmail] = useState("");
  const [deliveryAddress, setDeliveryAddress] = useState("");
  const [isPreCheckoutValid, setIsPreCheckoutValid] = useState(false);
  const { toast } = useToast();

  const handlePaymentSuccess = (response: any) => {
    console.log("Payment was successful!", response);
    // Handle order completion logic here
  };

  const handlePaymentClose = () => {
    console.log("Payment window was closed.");
  };

  const preCheckout = async () => {
    setIsLoading(true);
    setIsPreCheckoutValid(false);
    try {
      let accessToken = await getCookie("jeagereats_token");
      if (!accessToken?.value) {
        return toast({
          title: "You're not logged in",
          description: "You have to login in order to checkout your items",
          action: (
            <ToastAction altText="Login">
              <Link href={"/login"}>Login</Link>
            </ToastAction>
          ),
        });
      } else {
        let res = await axios.get(`${BASE_URL()}/users`, {
          headers: {
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
        setEmail(res?.data?.email);
        setIsPreCheckoutValid(true);
      }
    } catch (error) {
    } finally {
      setIsLoading(false);
    }
  };

  const handleCheckout = async () => {
    const placeOrder: PlaceOrderType = {
      restaurant_id: menu[0]?.restaurant_id || "",
      total_price: total + deliveryFee,
      delivery_fee: deliveryFee,
      delivery_address: deliveryAddress,
      items: menu.map((item) => ({
        menu_id: item.id,
        price_per_item: item.price,
        quantity: item.quantity,
      })),
    };

    try {
      let accessToken = await getCookie("jeagereats_token");
      if (!accessToken?.value) {
        return toast({
          title: "You're not logged in",
          description: "You have to login in order to checkout your items",
          action: (
            <ToastAction altText="Login">
              <Link href={"/login"}>Login</Link>
            </ToastAction>
          ),
        });
      } else {
        let res = await axios.post(`${BASE_URL()}/orders`, placeOrder, {
          headers: {
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
        if (res.status == 201) {
          resetCart();
          setIsOpen(false);
          router.push("/shop");
        }
        return toast({
          title: "Payment Successful!",
          description:
            "Your payment has been processed successfully. Thank you for your order!",
        });
      }
    } catch (error) {
      return toast({
        title: "Uh oh! Something went wrong.",
        description: "There was a problem with your request.",
      });
    }
  };

  const handleIncrement = (
    id: string,
    currentQuantity: number,
    stock: number
  ) => {
    if (currentQuantity < stock) {
      incrementQuantity(id);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="outline" size="icon" className="relative">
          <div className="flex items-center justify-center p-3 bg-primary text-white rounded-full">
            <FaShoppingCart />
          </div>
          {menu.length > 0 && (
            <span className="text-white absolute -top-2 -right-2 h-5 w-5 rounded-full bg-primary text-xs text-primary-foreground flex items-center justify-center">
              {menu.length}
            </span>
          )}
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Shopping Cart</DialogTitle>
        </DialogHeader>
        <ScrollArea className="max-h-[60vh]">
          <div className="space-y-4 pr-4">
            {menu.map((item) => (
              <Card key={item.id}>
                <CardContent className="p-4">
                  <div className="flex justify-between items-start gap-4">
                    <div className="space-y-1">
                      <h3 className="font-medium">{item.name}</h3>
                      <p className="text-sm text-muted-foreground">
                        {item.description}
                      </p>
                      <p className="font-medium">
                        ₦{item.price.toLocaleString()}
                      </p>
                    </div>
                    <Button
                      variant="ghost"
                      size="icon"
                      className="text-destructive"
                      onClick={() => removeFromMenu(item.id)}
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                  <div className="flex items-center gap-2 mt-4">
                    <Button
                      variant="outline"
                      size="icon"
                      className="h-8 w-8"
                      onClick={() => decrementQuantity(item.id)}
                    >
                      <Minus className="h-4 w-4" />
                    </Button>
                    <Input
                      type="number"
                      value={item.quantity}
                      readOnly
                      className="h-8 w-16 text-center"
                    />
                    <Button
                      variant="outline"
                      size="icon"
                      className="h-8 w-8"
                      onClick={() =>
                        handleIncrement(item.id, item.quantity, item.stock)
                      }
                    >
                      <Plus className="h-4 w-4" />
                    </Button>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </ScrollArea>
        <Separator className="my-4" />
        <div className="space-y-4">
          <div className="flex justify-between text-primary">
            <span className="font-medium">Subtotal</span>
            <span className="font-medium">₦{total.toLocaleString()}</span>
          </div>
          <div className="flex justify-between text-primary">
            <span className="font-medium">Delivery Fee</span>
            <span className="font-medium">₦{deliveryFee.toLocaleString()}</span>
          </div>
          <div className="flex justify-between text-primary">
            <span className="font-medium">Total</span>
            <span className="font-medium">
              ₦{(total + deliveryFee).toLocaleString()}
            </span>
          </div>
          <div className="mt-4">
            <Input
              placeholder="Enter delivery address"
              value={deliveryAddress}
              onChange={(e) => setDeliveryAddress(e.target.value)}
              className="w-full"
            />
          </div>
          {!isPreCheckoutValid ? (
            <Button
              className="w-full bg-primary text-white"
              size="lg"
              onClick={preCheckout}
              disabled={isLoading || total == 0 || !deliveryAddress}
            >
              {!isLoading ? "Checkout" : <Spinner bg="primary" />}
            </Button>
          ) : (
            <Button
              onClick={() => setIsOpen(false)}
              className=".paystack-payment-modal bg-primary text-white"
            >
              <PaymentButton
                email={email}
                amount={(total + deliveryFee) * 100}
                onSuccess={handleCheckout}
                onClose={handlePaymentClose}
              />
            </Button>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
}
