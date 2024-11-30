"use client";
import Image from "next/image";
import { ChangeEvent, FormEvent, useState } from "react";
import localPreview from "../../../public/assets/jollof.png";

const MenuAddForm = () => {
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
        <p>Upload Item Image</p>
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
      <div className="">
        <label htmlFor="">
          <p>Item Name</p>
          <input
            type="text"
            placeholder="Jollof Rice"
            className="outline w-full outline-none border border-primary p-2 mt-2 caret-primary rounded-lg"
          />
        </label>
      </div>
      <div>
        <label htmlFor="">
          <p>Item Description</p>
          <textarea
            name=""
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            id=""
          ></textarea>
        </label>
      </div>
      <div className="flex flex-col md:flex-row gap-6">
        <label htmlFor="">
          <p>Stock</p>
          <input
            type="number"
            placeholder="total item available"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
          />
        </label>
        <label htmlFor="">
          <p>Price</p>
          <input
            type="number"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
          />
        </label>
      </div>
      <button className="bg-primary text-white rounded-lg py-2">
        Add to Menu
      </button>
    </form>
  );
};

export default MenuAddForm;
