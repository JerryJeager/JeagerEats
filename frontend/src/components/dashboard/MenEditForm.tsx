"use client";
import Image from "next/image";
import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import axios from "axios";
import localPreview from "../../../public/assets/jollof.png";
import { BASE_URL } from "@/data";
import { getCookie } from "@/actions/handleCookies";
import { usePathname, useRouter } from "next/navigation";

interface FormData {
  name: string;
  description: string;
  price: number;
  stock: number;
}

const MenuEditForm = () => {
  let pathname = usePathname();
  let id = pathname.split("/").pop();
  const router = useRouter()
  const [profileImage, setProfileImage] = useState<null | File>(null);
  const [preview, setPreview] = useState<string | null>(null);
  const [formData, setFormData] = useState<FormData>({
    description: "",
    name: "",
    price: 0,
    stock: 0,
  });
  const [isLoading, setIsLoading] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

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

      // PUT request to update the menu item
      await axios.patch(`${BASE_URL()}/menus/${id}`, formData, {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken?.value}`,
        },
      });

      if (profileImage) {
        // Create FormData for the PATCH request to update the profile image
        const imageData = new FormData();
        imageData.append("file", profileImage);

        await axios.patch(`${BASE_URL()}/menus/img/${id}`, imageData, {
          headers: {
            "Content-Type": "multipart/form-data",
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
      }

      alert("Menu item updated successfully!");
    } catch (error) {
      console.error(error);
      alert("An error occurred. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!confirm("Are you sure you want to delete this menu item?")) return;
    setIsDeleting(true);

    try {
      let accessToken = await getCookie("jeagereats_token");

      // DELETE request to delete the menu item
      await axios.delete(`${BASE_URL()}/menus/${id}`, {
        headers: {
          Authorization: `Bearer ${accessToken?.value}`,
        },
      });

      alert("Menu item deleted successfully!");
      router.push("/dashboard/vendor/menu");
    } catch (error) {
      console.error(error);
      alert("An error occurred while deleting the menu item. Please try again.");
    } finally {
      setIsDeleting(false);
    }
  };

  useEffect(() => {
    const getMenuData = async () => {
      try {
        const res = await axios.get(`${BASE_URL()}/menus/${id}`);
        const data = res.data;
        setFormData({
          name: data.name,
          description: data.description,
          price: data.price,
          stock: data.stock,
        });
        if (data.img_url) {
          setPreview(data.img_url);
        } else {
          setPreview(localPreview as unknown as string);
        }
      } catch (error) {
        console.error("Error fetching menu data:", error);
      }
    };

    if (id) getMenuData();
  }, [id]);

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
        {isLoading ? "Submitting..." : "Edit"}
      </button>
      <button
        type="button"
        onClick={handleDelete}
        className="bg-red-500 text-white rounded-lg py-2"
        disabled={isDeleting}
      >
        {isDeleting ? "Processing..." : "Delete"}
      </button>
    </form>
  );
};

export default MenuEditForm;
