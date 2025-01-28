import { Restaurant, Role, User } from "@/types";
import { create } from "zustand";
import { persist } from "zustand/middleware";

interface Store {
  role: Role | ""; // Default value is an empty string
  user: User | null;
  restaurant: Restaurant | null;
  setUser: (user: User) => void;
  setRole: (role: Role) => void;
  setUserValue: (key: keyof User, value: any) => void;
  setRestaurant: (restaurant: Restaurant) => void;
  setRestaurantValue: (key: keyof Restaurant, value: any) => void;
  resetUser: () => void;
  resetRestaurant: () => void;
}

const useUserStore = create<Store>()(
  persist(
    (set) => ({
      role: "", 
      user: null,
      restaurant: null,
      setUser: (user: User) => set({ user }),
      setRole: (role: Role) => set({ role }),
      setUserValue: (key: keyof User, value: any) =>
        set((state) => ({
          user: {
            ...state.user,
            [key]: value,
          } as User, 
        })),
      setRestaurant: (restaurant: Restaurant) => set({ restaurant }),
      setRestaurantValue: (key: keyof Restaurant, value: any) =>
        set((state) => ({
          restaurant: state.restaurant ? {
            ...state.restaurant,
            [key]: value,
          } : { [key]: value } as Restaurant,
        })),
      resetUser: () => set({ user: null }),
      resetRestaurant: () => set({ restaurant: null }),
    }),
    {
      name: "user-store", 
      partialize: (state) => ({
        role: state.role,
        user: state.user,
      }), 
    }
  )
);

export default useUserStore;
