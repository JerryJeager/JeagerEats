"use client";
import Image from "next/image";
import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import axios from "axios";
import localPreview from "../../../public/assets/chef.png";
import { BASE_URL } from "@/data";
import { getCookie } from "@/actions/handleCookies";
import { Restaurant, RestaurantSelf } from "@/types";
import { useRouter } from "next/navigation";
import useUserStore from "@/store/useUserStore";
interface FormData {
  name: string;
  phone_number: string;
  description: string;
  address: string;
  opening_time: string;
  closing_time: string;
}

const VendorProfileForm = () => {
  const [profileImage, setProfileImage] = useState<null | File>(null);
  const [preview, setPreview] = useState<string | null>(null);

  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();
  const [restaurant, setRestaurant] = useState<Restaurant | null>(null);
  const formDataDefault: FormData = {
    name: "",
    phone_number: "",
    description: "",
    address: "",
    opening_time: "",
    closing_time: "",
  };
  const [formData, setFormData] = useState<FormData>({ ...formDataDefault });

  const { restaurant: restaurantStore, setRestaurant: setRestaurantStore } =
    useUserStore();

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    let { name, value } = e.target;
    setFormData((prev) => {
      return {
        ...prev,
        [name]: value,
      };
    });
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

      // Add a date to time fields to make them `time.Time` compatible
      const today = new Date().toISOString().split("T")[0]; // Get today's date in YYYY-MM-DD format
      const formattedData = {
        ...formData,
        opening_time: `${today}T${formData.opening_time}:00Z`, // Convert to ISO 8601
        closing_time: `${today}T${formData.closing_time}:00Z`,
      };

      // Submit normal field
      await axios.patch(`${BASE_URL()}/restaurants/profile`, formattedData, {
        headers: {
          Authorization: `Bearer ${accessToken?.value}`,
        },
      });

      // Upload profile image if available
      if (profileImage) {
        const imageData = new FormData();
        imageData.append("file", profileImage);

        await axios.patch(`${BASE_URL()}/restaurants/profile/img`, imageData, {
          headers: {
            "Content-Type": "multipart/form-data",
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
      }

      alert("Profile updated successfully!");
    } catch (error) {
      console.error("Error updating profile:", error);
      alert("Failed to update profile. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    const getRestaurant = async () => {
      try {
        const accessToken = await getCookie("jeagereats_token");
        const res = await axios.get(`${BASE_URL()}/restaurants/self`, {
          headers: {
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
        setRestaurant(res.data);
        setRestaurantStore(res.data);

        // Populate formData with fetched values
        setFormData({
          name: res.data.name || "",
          phone_number: res.data.phone_number || "",
          description: res.data.description || "",
          address: res.data.address || "",
          opening_time: res.data.opening_time?.slice(11, 16) || "", // Extract "HH:mm" from ISO string
          closing_time: res.data.closing_time?.slice(11, 16) || "", // Extract "HH:mm" from ISO string
        });

        // Set profile image preview if it exists and is not an empty string
        if (res.data.profile_img && res.data.profile_img.trim() !== "") {
          setPreview(res.data.profile_img);
        }
      } catch (error) {
        console.error("Error fetching restaurant data:", error);
        router.push("/auth/login");
      }
    };

    if (restaurantStore == null) {
      getRestaurant();
    } else {
      setRestaurant(restaurantStore);

      // Populate formData with store values if they exist
      setFormData({
        name: restaurantStore.name || "",
        phone_number: restaurantStore.phone_number || "",
        description: restaurantStore.description || "",
        address: restaurantStore.address || "",
        opening_time: restaurantStore.opening_time?.slice(11, 16) || "",
        closing_time: restaurantStore.closing_time?.slice(11, 16) || "",
      });

      // Set profile image preview if it exists and is not an empty string
      if (
        restaurantStore.profile_img &&
        restaurantStore.profile_img.trim() !== ""
      ) {
        setPreview(restaurantStore.profile_img);
      }
    }
  }, [restaurantStore]);

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-6 w-fit">
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
        <input type="file" onChange={handleChangeImage} accept="image/*" />
      </div>

      <div className="flex flex-col md:flex-row gap-6">
        <label>
          <p>Restaurant Name</p>
          <input
            name="name"
            type="text"
            placeholder="Chitis"
            className="outline w-full outline-none border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.name}
            onChange={handleChange}
          />
        </label>
        <label>
          <p>Restaurant Phone</p>
          <input
            name="phone_number"
            type="text"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.phone_number}
            onChange={handleChange}
          />
        </label>
      </div>

      <div>
        <label>
          <p>Restaurant Description</p>
          <textarea
            name="description"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.description}
            onChange={handleChange}
          ></textarea>
        </label>
      </div>

      <div>
        <label>
          <p>Restaurant Address</p>
          <input
            name="address"
            type="text"
            className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
            value={formData.address}
            onChange={handleChange}
          />
        </label>
      </div>

      <div className="flex gap-8">
        <label className="w-full">
          <p>Restaurant Opening Time</p>
          <input
            name="opening_time"
            type="time"
            className="outline-none border border-primary p-2 mt-2 caret-primary w-full rounded-lg"
            value={formData.opening_time}
            onChange={handleChange}
          />
        </label>
        <label className="w-full">
          <p>Restaurant Closing Time</p>
          <input
            name="closing_time"
            type="time"
            className="outline-none border border-primary p-2 mt-2 caret-primary w-full rounded-lg"
            value={formData.closing_time}
            onChange={handleChange}
          />
        </label>
      </div>

      <button
        type="submit"
        className="bg-primary text-white rounded-lg py-2"
        disabled={isLoading}
      >
        {isLoading ? "Saving Changes..." : "Save Changes"}
      </button>
    </form>
  );
};

export default VendorProfileForm;
