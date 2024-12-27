// src/stores/CategoryStore.js
import { writable } from 'svelte/store';
import { Clapperboard, Hexagon, TrendingUp, Shirt, HeartPulse, House, Car, Hospital, Pizza } from 'lucide-svelte';


// Initial categories data
const defaultCategories = [
  {
    name: 'Food',
    icon: Pizza,
    subCategories: ['Groceries', 'Restaurants', 'Snacks', 'Takeout']
  },
  {
    name: "Housing",
    icon: House,
    subCategories: ["", "",]
  },
  {
    name: 'Transportation',
    icon: Car,
    subCategories: ['Car', 'Public Transport', 'Fuel', 'Parking']
  },
  {
    name: "Healthcare",
    icon: HeartPulse,
    subCategories: ["", "", "",]
  },
  {
    name: "Miscellaneous",
    icon: Hexagon,
    subCategories: ["", "", "",]
  },
  {
    name: "Clothing",
    icon: Shirt,
    subCategories: ["", "", "",]
  },
  {
    name: 'Income',
    icon: TrendingUp,
    subCategories: ['Salary', 'Investments', 'Gifts', 'Other Income']
  },
  {
    name: "Entertainment",
    icon: Clapperboard,
    subCategories: ["", "",]
  },
  {
    name: 'Health',
    icon: Hospital,
    subCategories: ['Insurance', 'Doctors', 'Medications', 'Gym']
  }
];

// Create the writable store
const categoriesStore = writable(defaultCategories);

// Utility functions
const initializeCategories = () => {
  categoriesStore.set(defaultCategories); // Reset to default categories
};

const addCategory = (category) => {
  categoriesStore.update((current) => [...current, category]);
};

const assignIconToCategory = (categoryName, icon) => {
  categoriesStore.update((current) =>
    current.map((category) =>
      category.name === categoryName ? { ...category, icon } : category
    )
  );
};

const getCategoryIcon = (categoryName) => {
  let icon = null;
  categoriesStore.subscribe((current) => {
    const category = current.find((cat) => cat.name === categoryName);
    icon = category ? category.icon : null;
  })();
  return icon;
};

// Export the store and helper functions
export const categories = {
  subscribe: categoriesStore.subscribe,
  initializeCategories,
  addCategory,
  assignIconToCategory,
  getCategoryIcon
};
