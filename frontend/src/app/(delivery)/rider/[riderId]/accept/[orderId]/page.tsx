"use client"

import { BASE_URL } from "@/data"
import axios from "axios"
import { usePathname, useRouter } from "next/navigation"
import {useEffect, useState} from "react"
import woman from "../../../../../../../public/assets/womanrider.svg"
import Image from "next/image"

const AcceptOrder = () => {
  const pathname = usePathname()
  const router = useRouter()
  const [isLoading, setIsLoading] = useState(true)
  useEffect(() => {
    const acceptDeliveryReq = async () => {
        setIsLoading(true)
        try{
            const reqBody = {
                rider_id: pathname?.split("/")[2]
            }
            let res = await axios.post(`${BASE_URL()}/orders/${pathname?.split("/")[4]}/accept`, reqBody)

            if (res.status == 201){
                console.log("successful")
            }
        }catch{
            router.push("/rejected")
        }finally{
            setIsLoading(false)
        }
    }
    acceptDeliveryReq()
  }, [])
  return (
    <>
  {isLoading && <div className="flex items-center justify-center min-h-screen">Loading...</div>}

  {!isLoading && (
    <div className="flex items-center justify-center min-h-screen bg-gray-50">
      <div className="bg-white rounded-lg shadow-md p-6 max-w-md text-center">
        {/* Success Icon/Image */}
        <div className="mb-6">
          <Image
            src={woman} 
            alt="Success"
            className="w-32 h-32 mx-auto object-cover"
          />
        </div>

        {/* Confirmation Message */}
        <h1 className="text-xl font-semibold text-green-600 mb-4">
          Delivery Request Accepted!
        </h1>
        <p className="text-gray-700 mb-6">
          You've successfully accepted the delivery request. Please head to the restaurant to pick up the package.
        </p>

        
      </div>
    </div>
  )}
</>

  )
}

export default AcceptOrder