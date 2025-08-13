import useAuthStore from "@/stores/authStore";
import { roles } from "@/utils/const";
import ChefHomeScreen from "./chefHome";
import UserHomeScreen from "./userHome";
import AdminHomeScreen from "./adminHome";
import NotFoundScreen from "@/screens/notFound";
import Navbar from "@/components/navbar";
import { useEffect } from "react";
import VerifySignedIn from "@/utils/verify";
import { Toaster } from "react-hot-toast";

const roleScreens = {
    [roles.CHEF]: <ChefHomeScreen />,
    [roles.CUSTOMER]: <UserHomeScreen />,
    [roles.ADMIN]: <AdminHomeScreen />,
}


export default function HomeScreen() {
    const { role } = useAuthStore.getState();

    useEffect(() => {VerifySignedIn()} , [])

    let homeScreen = roleScreens[role] || <NotFoundScreen />;
    return <div className="h-screen w-screen flex flex-col">
        <Navbar/>
        <Toaster/>
        <title>Home - InOrder</title>
        {homeScreen}
    </div>
}