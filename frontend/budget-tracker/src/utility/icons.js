import * as Icons from 'lucide-svelte'; // Assuming you're using lucide-svelte

export const iconComponents = {
  Transportation: Icons.Car,
  Entertainment: Icons.Clapperboard,
  Food: Icons.Pizza,
  Housing: Icons.House,
  Healthcare: Icons.HeartPulse,
  Miscellaneous: Icons.Hexagon,
  Clothing: Icons.Shirt,
  Utilities: Icons.Cable,
  Savings: Icons.PiggyBank,
  Education: Icons.BookOpen,
  Insurance: Icons.ShieldCheck,
  Pets: Icons.PawPrint,
  Travel: Icons.Plane,
  Gifts: Icons.Gift,
  Subscriptions: Icons.FileText,
  PersonalCare: Icons.Smile,
  DebtPayments: Icons.CreditCard,
  Salary: Icons.CreditCard,
  Shopping: Icons.ShoppingBag,

  ArrowUp: Icons.ChevronUp,
  ArrowDown: Icons.ChevronDown
};


export function getIconComponent(iconName) {
  return iconComponents[iconName] || null;
}