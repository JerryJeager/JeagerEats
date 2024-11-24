import Image from "next/image"
import logo from "../../public/assets/logo.png"
import Link from "next/link"
const Header = () => {
  return (
    <header className="px-[5%] lg:px-[8%] py-4 lg:py-6 flex items-center justify-between bg-white">
        <div>
            <Link href={"/"}> <Image src={logo} placeholder="blur" alt="logo" width={140} /></Link>
        </div>
        <Link href={"auth/login"}>
            <button className="bg-primary py-2 px-8 rounded-lg text-white">Login</button>
        </Link>
    </header>
  )
}

export default Header