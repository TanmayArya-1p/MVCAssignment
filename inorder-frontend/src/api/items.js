
export async function getAllItems(limit,offset) {
    try {
        const response = await axios.get(`/api/items?limit=${limit}&offset=${offset}`);
        return response.data;
    } catch (error) {
        console.error("Error fetching items:", error);
        throw error;
    }
}