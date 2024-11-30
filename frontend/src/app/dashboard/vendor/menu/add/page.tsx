import MenuAddForm from "@/components/dashboard/MenuAddForm";

const MenuAdd = () => {
  return (
    <section className="mt-8">
      <h2 className="font-bold text-2xl md:text-3xl text-center">
        Add Menu Item
      </h2>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg bg-white  h-full md:text-lg">
        <MenuAddForm />
      </div>
    </section>
  );
};

export default MenuAdd;
