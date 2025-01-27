import { RestaurantMenuCardType } from "@/types";
import { create } from "zustand";
import { persist } from "zustand/middleware";

interface Store {
  menu: RestaurantMenuCardType[];
  addToMenu: (menuItem: RestaurantMenuCardType) => void;
  removeFromMenu: (id: string) => void;
  resetCart: () => void;
  incrementQuantity: (id: string) => void;
  decrementQuantity: (id: string) => void;
}

const useCartStore = create<Store>()(
  persist(
    (set) => ({
      menu: [],

      addToMenu: (menuItem: RestaurantMenuCardType) =>
        set((state) => {
          const existingItem = state.menu.find((item) => item.id === menuItem.id);
          if (existingItem) {
            // If the item already exists, increment its quantity
            return {
              menu: state.menu.map((item) =>
                item.id === menuItem.id
                  ? { ...item, quantity: item.quantity + 1 }
                  : item
              ),
            };
          } else {
            // If the item is new, add it to the cart with quantity set to 1
            return {
              menu: [...state.menu, { ...menuItem, quantity: 1 }],
            };
          }
        }),

      removeFromMenu: (id: string) =>
        set((state) => ({
          menu: state.menu.filter((item) => item.id !== id),
        })),

      resetCart: () =>
        set(() => ({
          menu: [],
        })),

      incrementQuantity: (id: string) =>
        set((state) => ({
          menu: state.menu.map((item) =>
            item.id === id ? { ...item, quantity: item.quantity + 1 } : item
          ),
        })),

      decrementQuantity: (id: string) =>
        set((state) => ({
          menu: state.menu.map((item) =>
            item.id === id && item.quantity > 1
              ? { ...item, quantity: item.quantity - 1 }
              : item
          ),
        })),
    }),

    {
      name: "cart-store",
      partialize: (state) => ({
        menu: state.menu,
      }),
    }
  )
);

export default useCartStore;
