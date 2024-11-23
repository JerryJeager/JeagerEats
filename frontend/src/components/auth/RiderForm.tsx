"use client";
import { BASE_URL, BASE_URL_LOCAL } from "@/data";
import axios from "axios";
import {
  ChangeEvent,
  Dispatch,
  FormEvent,
  SetStateAction,
  useState,
} from "react";
import Spinner from "../ui/Spinner";
import { Role } from "@/types";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { useRouter } from "next/navigation";

const role: Role = { name: "rider" };

interface FormData {
  first_name: string;
  last_name: string;
  email: string;
  address: string;
  password: string;
  phone_number: string;
  role: string;
}

const RiderForm = () => {
  const router = useRouter();
  const formDataDefault: FormData = {
    first_name: "",
    address: "",
    email: "",
    last_name: "",
    password: "",
    phone_number: "",
    role: role.name,
  };
  const [formData, setFormData] = useState<FormData>({ ...formDataDefault });
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isPasswordShown, setIsPasswordShown] = useState(false);
  const [isConfirmPasswordShown, setIsConfirmPasswordShown] = useState(false);
  const handleTogglePasswordShown = (
    setP: Dispatch<SetStateAction<boolean>>
  ) => {
    setP((prev) => !prev);
  };
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    let { name, value } = e.target;
    setFormData((prev) => {
      return {
        ...prev,
        [name]: value,
      };
    });
  };
  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");
    setFormData((prev) => {
      const trimmedData: FormData = { ...formDataDefault };
      for (const key in prev) {
        trimmedData[key as keyof FormData] =
          prev[key as keyof FormData]?.trim();
      }
      return trimmedData;
    });

    if (formData.password !== confirmPassword) {
      setError("password mismatch");
      setIsLoading(false);
      return;
    }

    try {
      const res = await axios.post(`${BASE_URL}/users/signup`, formData);
      console.log("User signed up successfully:", res.data);
      router.push("auth/login")
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response) {
          setError(error?.response?.data?.message);
          //   console.log(error?.response?.data?.message);
          //   console.error("Error Response:", error.response.data);
          //   console.error("Status Code:", error.response.status);
          setError(error.response.data?.message);
        }
        // } else if (error.request) {
        //   console.error("Error Request:", error.request);
        // }
        else {
          //   console.error("Error Message:", error.message);
          setError("Network Error");
        }
      }
    } finally {
      setIsLoading(false);
    }
  };
  return (
    <form
      onSubmit={handleSubmit}
      className="rounded-2xl w-full md:w-1/2 bg-white shadow-xl p-6 vendor-form"
    >
      <div className="flex flex-col gap-4 md:gap-6">
        <div className="flex flex-col md:flex-row gap-4">
          <label htmlFor="" className="w-full">
            <p>First Name</p>
            <input
              name="first_name"
              value={formData.first_name}
              onChange={handleChange}
              type="text"
              placeholder="John"
              required
            />
          </label>
          <label htmlFor="" className="w-full">
            <p>Last Name</p>
            <input
              value={formData.last_name}
              name="last_name"
              onChange={handleChange}
              type="text"
              placeholder="Doe"
              required
            />
          </label>
        </div>
        <div className="flex flex-col md:flex-row gap-4">
          <label htmlFor="" className="w-full">
            <p>Email</p>
            <input
              value={formData.email}
              name="email"
              onChange={handleChange}
              type="email"
              placeholder="johndoe@gmail.com"
              required
            />
          </label>
          <label htmlFor="" className="w-full">
            <p>Phone Number</p>
            <input
              name="phone_number"
              onChange={handleChange}
              value={formData.phone_number}
              type="text"
              required
              placeholder="07082562858"
            />
          </label>
        </div>
        <label htmlFor="">
          <p>Address</p>
          <input
            value={formData.address}
            name="address"
            onChange={handleChange}
            type="text"
            placeholder="house address"
            required
          />
        </label>
        <label htmlFor="">
          <p>Password</p>
          <div className="relative">
            <input
              name="password"
              value={formData.password}
              onChange={handleChange}
              type={isPasswordShown ? "text" : "password"}
              required
              placeholder="password"
            />
            <button
              type="button"
              onClick={() => handleTogglePasswordShown(setIsPasswordShown)}
              className="absolute right-2 md:right-3 bottom-3"
            >
              {isPasswordShown ? <LuEye /> : <LuEyeOff />}
            </button>
          </div>
        </label>
        <label htmlFor="">
          <p>Confirm Password</p>
          <div className="relative">
            <input
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              name="confirm_password"
              type={isConfirmPasswordShown ? "text" : "password"}
              placeholder="password"
            />
            <button
              type="button"
              onClick={() =>
                handleTogglePasswordShown(setIsConfirmPasswordShown)
              }
              className="absolute right-2 md:right-3 bottom-3"
            >
              {isConfirmPasswordShown ? <LuEye /> : <LuEyeOff />}
            </button>
          </div>
        </label>
        {error && <p className="text-red-500">{error}</p>}
        <button className="w-full rounded-lg text-white font-semibold text-center bg-primary p-2 md:p-3">
          {!isLoading ? "Register" : <Spinner bg="primary" />}
        </button>
      </div>
    </form>
  );
};

export default RiderForm;
