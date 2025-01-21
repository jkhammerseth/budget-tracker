import { categories } from "../../stores/categories";

export async function getCategories() {
    try {
        const category_response = await fetch('http://localhost:8080/api/users/categories',{
                credentials: 'include'
        });
        if (!category_response.ok) {
            throw new error(`HTTP error! status: ${category_response.status}`);
        }

        const category_data =  await category_response.json();
        categories.set(category_data);
    } catch (error) {
        console.error("Failed to fetch categories:", error);
    }
}

export async function getCategoryById(category) {
    try {
        const response = await fetch(`http://localhost:8080/api/users/category/${category.id}`, { 
            credentials: "include" 
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || "Failed to get category");
        }

        const data = await response.json();
        return data; 
    } catch (error) {
        console.error("Error fetching category:", error);
        throw error;
    }
}
