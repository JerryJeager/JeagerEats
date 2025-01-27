"use client"
import React, { useEffect, useState } from "react";
import dynamic from 'next/dynamic';

// Dynamically import PaystackButton to disable SSR
const PaystackButton = dynamic(
  () => import('react-paystack').then((mod) => mod.PaystackButton),
  { ssr: false }
);

interface PaymentButtonProps {
  email: string;
  amount: number;
  onSuccess: (response: any) => void;
  onClose: () => void;
}

const PaymentButton: React.FC<PaymentButtonProps> = ({
  email,
  amount,
  onSuccess,
  onClose,
}) => {
  const publicKey = "pk_test_ede30c5c22a7d2c64ae47446114c3464ac618a5a";
  const [callbackUrl, setCallbackUrl] = useState("");

  useEffect(() => {
    // Client-side check to ensure window is available
    if (typeof window !== 'undefined') {
      setCallbackUrl(window.location.origin + "/shop");
    }
  }, []);

  const componentProps = {
    email,
    amount,
    publicKey,
    metadata: {
      custom_fields: [
        {
          display_name: "Email",
          variable_name: "email",
          value: email,
        },
      ],
    },
    text: "Pay Now",
    onSuccess: (response: any) => {
      console.log("Payment successful", response);
      onSuccess(response);
    },
    onClose: () => {
      console.log("Payment closed");
      onClose();
    },
    callback_url: callbackUrl,
  };

  return <PaystackButton {...componentProps} />;
};

export default PaymentButton;