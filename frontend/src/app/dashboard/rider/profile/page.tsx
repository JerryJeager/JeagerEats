"use client";
import RiderProfileForm from "@/components/dashboard/rider/RiderProfileForm";

const Profile = () => {
  return (
    <section className="mt-8">
      <h2 className="font-bold text-2xl md:text-3xl text-center">
        Profile Management
      </h2>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg bg-white  h-full md:text-lg ">
       <RiderProfileForm />
      </div>
    </section>
  );
};

export default Profile;
