"use client";
import Image from "next/image";
import { ChangeEvent, FormEvent, useState } from "react";
import axios from "axios";
import localPreview from "../../../public/assets/jollof.png";
import { BASE_URL } from "@/data";
import { getCookie } from "@/actions/handleCookies";

interface FormData {
  name: string;
  description: string;
  price: number;
  stock: number;
}

const MenuAddForm = () => {
  const [profileImage, setProfileImage] = useState<null | File>(null);
  const [preview, setPreview] = useState<string | null>(null);
  const [formData, setFormData] = useState<FormData>({
    description: "",
    name: "",
    price: 0,
    stock: 0,
  });
  const [isLoading, setIsLoading] = useState(false);

  const handleChange = (
    e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]:
        name === "price" || name === "stock" ? parseFloat(value) || 0 : value,
    }));
  };

  const handleChangeImage = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setProfileImage(file);

      // Create a preview of the image
      const reader = new FileReader();
      reader.onload = () => {
        setPreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      let accessToken = await getCookie("jeagereats_token");
      // POST request to create a new menu item
      const response = await axios.post(`${BASE_URL()}/menus`, formData, {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken?.value}`,
        },
      });

      const menuId = response.data.id;

      if (profileImage && menuId) {
        // Create FormData for the PATCH request to update the profile image
        const formData = new FormData();
        formData.append("file", profileImage); // Changed the key name to "file"

        await axios.patch(`${BASE_URL()}/menus/img/${menuId}`, formData, {
          headers: {
            "Content-Type": "multipart/form-data",
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
      }

      alert("Menu item added successfully!");
    } catch (error) {
      console.error(error);
      alert("An error occurred. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-6 w-fit">
      <div>
        <p>Upload Item Image</p>
        <Image
          src={preview || localPreview}
          className="rounded-full w-[200px] h-[200px] my-8 mx-auto"
          alt="image-preview"
          objectFit="cover"
          width={200}
          height={200}
        />
        <input
          type="file"
          onChange={handleChangeImage}
          accept="image/*"
          required
        />
      </div>
      <div>
        <label>
          <p>Item Name</p>
          <input
            type="text"
            name="name"
            placeholder="Jollof Rice"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.name}
            onChange={handleChange}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <p>Item Description</p>
          <textarea
            name="description"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.description}
            onChange={handleChange}
            required
          ></textarea>
        </label>
      </div>
      <div className="flex flex-col md:flex-row gap-6">
        <label>
          <p>Stock</p>
          <input
            type="number"
            name="stock"
            placeholder="Total items available"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.stock}
            onChange={handleChange}
            required
          />
        </label>
        <label>
          <p>Price</p>
          <input
            type="number"
            name="price"
            placeholder="Price"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.price}
            onChange={handleChange}
            required
          />
        </label>
      </div>
      <button
        type="submit"
        className="bg-primary text-white rounded-lg py-2"
        disabled={isLoading}
      >
        {isLoading ? "Submitting..." : "Add to Menu"}
      </button>
    </form>
  );
};

export default MenuAddForm;
