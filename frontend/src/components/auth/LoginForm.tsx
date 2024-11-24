"use client";
import { BASE_URL } from "@/data";
import axios from "axios";
import {
  ChangeEvent,
  Dispatch,
  FormEvent,
  SetStateAction,
  useState,
} from "react";
import Spinner from "../ui/Spinner";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { useRouter } from "next/navigation";
import { storeCookie } from "@/actions/handleCookies";
import { Role, Roles } from "@/types";

interface FormData {
  email: string;
  password: string;
}

const LoginForm = () => {
  const router = useRouter();
  const formDataDefault: FormData = {
    email: "",
    password: "",
  };
  const [formData, setFormData] = useState<FormData>({ ...formDataDefault });
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isPasswordShown, setIsPasswordShown] = useState(false);
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

    try {
      const res = await axios.post(`${BASE_URL()}/users/login`, formData);
      console.log("User logged in successfully:", res.data);
      if (res.status == 200) {
        await storeCookie("jeagereats_user_id", res.data?.id);
        await storeCookie("jeagereats_token", res.data?.token);
        const role: Roles = res.data?.role as Roles;
        role == "rider"
          ? router.push("/dashboard/rider")
          : role == "vendor"
          ? router.push("/dashboard/vendor")
          : router.push("/shop");
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response) {
          setError(error?.response?.data?.message);
        } else {
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
      className="rounded-2xl w-full h-fit md:w-1/2 bg-white shadow-xl p-6 vendor-form"
    >
      <div className="flex flex-col gap-4 md:gap-6">
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
        {error && <p className="text-red-500">{error}</p>}
        <button className="w-full rounded-lg text-white font-semibold text-center bg-primary p-2 md:p-3">
          {!isLoading ? "Login" : <Spinner bg="primary" />}
        </button>
      </div>
    </form>
  );
};

export default LoginForm;
