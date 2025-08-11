import useAuthStore from "../../stores/authStore";
import { roles } from "../../utils/const";
import ChefHomeScreen from "./chef-home";
import UserHomeScreen from "./user-home";
import AdminHomeScreen from "./admin-home";
import NotFoundScreen from "../../screens/not-found";
import Navbar from "../../components/navbar";
import { useEffect } from "react";
import VerifySignedIn from "../../utils/verify";

const roleScreens = {
    [roles.CHEF]: <ChefHomeScreen />,
    [roles.CUSTOMER]: <UserHomeScreen />,
    [roles.ADMIN]: <AdminHomeScreen />,
}


export default function HomeScreen() {
    const { role } = useAuthStore.getState();

    useEffect(() => {VerifySignedIn()} , [])

    let homeScreen = roleScreens[role] || <NotFoundScreen />;
    return <>
        <Navbar role={role} />
        {homeScreen}
    </>
}