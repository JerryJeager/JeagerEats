"use client"
import { FormEvent } from "react";

const VendorForm = () => {
  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
  };
  return (
    <form
      onSubmit={handleSubmit}
      className="rounded-2xl w-full md:w-1/2 bg-white shadow-xl p-6 vendor-form"
    >
      <div className="flex flex-col gap-4 md:gap-8">
        <label htmlFor="">
          <p>First Name</p>
          <input type="text" placeholder="John" />
        </label>
        <label htmlFor="">
          <p>Last Name</p>
          <input type="text" placeholder="Doe" />
        </label>
        <label htmlFor="">
          <p>Email</p>
          <input type="email" placeholder="johndoe@gmail.com" />
        </label>
        <label htmlFor="">
          <p>Phone Number</p>
          <input type="text" placeholder="07082562858" />
        </label>
        <label htmlFor="">
          <p>Address</p>
          <input type="text" placeholder="house address" />
        </label>
        <button className="w-full mt-3 md:mt-4 rounded-lg text-white font-semibold text-center bg-primary p-2 md:p-3">
          Register
        </button>
      </div>
    </form>
  );
};

export default VendorForm;
