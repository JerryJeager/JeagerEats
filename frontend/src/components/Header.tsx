import Image from "next/image"
import logo from "../../public/assets/logo.png"
import Link from "next/link"
const Header = () => {
  return (
    <header className="px-[5%] lg:px-[10%] py-4 lg:py-6 flex items-center justify-between">
        <div>
            <Link href={"/"}> <Image src={logo} placeholder="blur" alt="logo" width={140} /></Link>
        </div>
        <Link href={"/"}>
            <button className="bg-primary py-2 px-8 rounded-lg text-white">Login</button>
        </Link>
    </header>
  )
}

export default Header