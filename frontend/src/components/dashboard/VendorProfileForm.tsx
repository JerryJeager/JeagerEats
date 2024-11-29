"use client";
import Image from "next/image";
import { ChangeEvent, FormEvent, useState } from "react";
import localPreview from "../../../public/assets/chef.png";

const VendorProfileForm = () => {
  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
  };
  const [profileImage, setProfileImage] = useState<null | any>();
  const [preview, setPreview] = useState<any>(null);
  const handleChangeImage = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setProfileImage(file);

      // Create a preview of the image
      const reader = new FileReader();
      reader.onload = () => {
        setPreview(reader.result);
      };
      reader.readAsDataURL(file);
    }
  };
  return (
    <form
      action=""
      onSubmit={handleSubmit}
      className="flex flex-col gap-6 w-fit "
    >
      <div>
        <p>Upload Restaurant Image/logo</p>
        <Image
          src={preview ? preview : localPreview}
          className="rounded-full w-[200px] h-[200px] my-8 mx-auto"
          alt="image-preview"
          objectFit="cover"
          width={200}
          height={200}
        />
        <input
          type="file"
          src=""
          onChange={handleChangeImage}
          alt=""
          accept="image/*"
        />
      </div>
      <div className="flex flex-col md:flex-row gap-6">
        <label htmlFor="">
          <p>Restaurant Name</p>
          <input
            type="text"
            placeholder="Chitis"
            className="outline w-full outline-none border border-primary p-2 mt-2 caret-primary rounded-lg"
          />
        </label>
        <label htmlFor="">
          <p>Restaurant Phone</p>
          <input
            type="text"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
          />
        </label>
      </div>
      <div>
        <label htmlFor="">
          <p>Restaurant Description</p>
          <textarea
            name=""
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            id=""
          ></textarea>
        </label>
      </div>
      <div>
        <label htmlFor="">
          <p>Restaurant Address</p>
          <input
            type="text"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
          />
        </label>
      </div>
      <div className="flex gap-8">
        <label htmlFor="" className="w-full">
          <p>Restaurant Opening Time</p>
          <input
            type="time"
            className="outline-none border border-primary p-2 mt-2 caret-primary w-full rounded-lg"
          />
        </label>
        <label htmlFor="" className="w-full">
          <p>Restaurant Closing Time</p>
          <input
            type="time"
            className="outline-none border border-primary p-2 mt-2 caret-primary w-full  rounded-lg"
          />
        </label>
      </div>
      <button className="bg-primary text-white rounded-lg py-2">
        Save Changes
      </button>
    </form>
  );
};

export default VendorProfileForm;
