import { pacifico } from "./Hero"

const Footer = () => {
  return (
    <footer className="bg-primary py-8 px-[5%] lg:px-[8%]">
      <h2 className={`${pacifico.className} text-xs font-semibold text-white`}>&copy; 2025, JeagerEats. All rights reserved</h2>
    </footer>
  )
}

export default Footer