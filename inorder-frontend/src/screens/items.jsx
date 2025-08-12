import { useEffect, useState } from "react";
import useAuthStore from "../stores/authStore";
import ItemMenu from "../components/itemMenu";
import Navbar from "../components/navbar";
import CreateItemForm from "../components/createItemForm";
import Spinner from "../components/spinner";
import VerifySignedIn from "../utils/verify";
import { useNavigate } from "react-router-dom";
import { Toaster } from "react-hot-toast";


export default function ItemScreen() {
    const {role,username} = useAuthStore.getState();
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate();
    useEffect(() => {VerifySignedIn()}, [])

    useEffect(() => {
        if (role !== "admin") {
            navigate("/notfound");
        } else {
            setLoading(false);
        }
    },[])

    if (loading) {
        return <div className="h-screen w-screen flex items-center justify-center">
            <Spinner />
        </div>
    }


    return <div className="h-screen w-screen flex flex-col">
        <title>Items - InOrder</title>
        <Navbar></Navbar>
        <Toaster></Toaster>
        <div className="flex flex-col p-5 items-center w-full">
            <ItemMenu admin></ItemMenu>
            <div className="mt-10">
                <div className="text-3xl font-bold mb-2">Create Item</div>
                <CreateItemForm />
            </div>
        </div>
    </div>
}