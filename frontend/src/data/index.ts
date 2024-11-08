import bike from "../../public/assets/Motorcycle.png"
import shop from "../../public/assets/shop.png"
import cutlery from "../../public/assets/cutlery.png"
import { StaticImageData } from "next/image"

export type Join = {
    title: string 
    content: string 
    cta: string 
    link: string 
    icon: StaticImageData
    color: string 
}

export const JoinData: Join[] = [
    {
        title: "Earn with Flexibility", 
        content: "Deliver with us and earn whenever you want. Flexibility meets opportunity", 
        cta: "Sign up as a Rider",
        link: "", 
        icon: bike,
        color: "#FFE5CC"
    },
    {
        title: "Grow your Restaurant", 
        content: "Reach more customers and boost your sales. Partner with JaegerEats", 
        cta: "Become a Partner",
        link: "", 
        icon: shop,
        color: "#F0EDEE"
    },
    {
        title: "Order in Seconds", 
        content: "Get your favorite meals delivered fast. Simple, easy, and reliable.", 
        cta: "Order Now",
        link: "", 
        icon: cutlery, 
        color: "#CDE7FF"
    },
]