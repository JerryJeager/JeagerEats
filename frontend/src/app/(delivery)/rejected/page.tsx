import Image from "next/image";
import missed from "../../../../public/assets/missedchance.svg"


const RejectedOrder = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white rounded-lg shadow-md p-6 max-w-md text-center">
        {/* Image */}
        <div className="mb-6">
          <Image
            src={missed} // Replace with the actual image path
            alt="Order Already Taken"
            className="w-32 h-32 mx-auto object-cover"
          />
        </div>

        {/* Message */}
        <h1 className="text-xl font-semibold text-red-500 mb-4">
          Sorry, this delivery has already been taken!
        </h1>
        <p className="text-gray-600 mb-6">
          It seems another rider has accepted this order. Check the available
          deliveries on your email and grab the next one!
        </p>

        {/* Action Button */}
        
      </div>
    </div>
  );
};

export default RejectedOrder;
