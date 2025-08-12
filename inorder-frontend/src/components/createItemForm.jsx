
import { uploadImage, createItem } from "../api/items";
import {toast,Toaster} from "react-hot-toast";

export default function CreateItemForm() {


    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const image = formData.get("image");
        try {
            console.log("Creating Item",image)
            let resp = null
            if(image) {
                resp = await uploadImage(image);
                console.log(resp)
            }
            const createResp = await createItem({
                name: formData.get("name"),
                price: formData.get("price"),
                description: formData.get("description"),
                tags: formData.get("item-tags").split(",").map(tag => tag.trim()),
                image: resp ? resp.url : null,
            });
            console.log("creating",createResp);
            window.location.reload()
        } catch(error) {
            console.error("error creating item:", error);
            toast.error("Failed to create item");
        }

    }

    return <div id="create-item-container" className="flex flex-col gap-4">
            <Toaster></Toaster>
            <form onSubmit={handleSubmit} id="create-item-form" className="flex flex-col gap-3 ubuntu-regular bg-white border-2 p-3 rounded" style={{width:"30rem"}}>
                <div className="flex flex-row justify-between">
                    <div>Name:</div>
                    <input type="text" id="name" name="name" className="border-2 rounded-sm p-1 px-2"required placeholder="Burger"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Description:</div>
                    <input type="text" id="description" name="description" className="border-2 rounded-sm p-1 px-2" required placeholder="just a burger"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Price:</div>
                    <input type="text" id="price" name="price" className="border-2 rounded-sm p-1 px-2" required placeholder="30"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Tags:</div>
                    <input type="text" id="item-tags" name="item-tags" className="border-2 rounded-sm p-1 px-2" required placeholder="spicy,veg"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Image:</div>
                    <input type="file" id="file" name="image" accept="image/*" className="file-input"/>
                </div>
                <button type="submit" className="sbutton ubuntu-bold">Create Item</button>
            </form>
        </div>
}