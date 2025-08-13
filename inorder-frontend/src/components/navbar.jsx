import * as auth from "../api/auth";
import useAuthStore from "../stores/authStore";
import { useNavigate } from "react-router-dom";
import { roles } from "../utils/const";
import MenuIcon from "../icons/menuIcon";
import LogoutIcon from "../icons/logoutIcon";


export default function Navbar() {
    const { role ,username} = useAuthStore.getState();

    const navigate = useNavigate();

    const logoutHandler = async () => {
        if(!window.confirm("Are you sure you want to logout?")) {
            return
        }
        const {setUsername, setAuthToken, setRefreshToken,setRole, setUserID} = useAuthStore.getState();

        setUsername("");
        setAuthToken("");
        setRefreshToken("");
        setRole("");

        await auth.LogoutUser();
        navigate("/");
    }



    return <>
        <nav className="bg-white border-b-2 border-black">
        <div className="max-w-screen flex flex-wrap items-center justify-between mx-auto p-4">
            <a href="/home" className="flex items-center space-x-3 rtl:space-x-reverse">
                <span className="self-center text-3xl ubuntu-bold">InOrder</span>
                <span className="flex items-end text-lg ubuntu-regular self-end">{role == roles.CUSTOMER ? "" : role}</span>
            </a>
            <div className="flex flex-row gap-4 text-lg">
                    {role == roles.ADMIN && <>
                        <li className="items-center align-middle flex">
                            <a className='order-link' href="/items">Items</a>
                        </li>
                        <li className="items-center align-middle flex">
                            <a className='order-link' href="/users">Users</a>
                        </li>
                    </>}
            </div>
            <button data-collapse-toggle="navbar-default" type="button" className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200" aria-controls="navbar-default" aria-expanded="false">
                <span className="sr-only">Open main menu</span>
                <MenuIcon className="w-6 h-6" />
            </button>
            <div className="hidden w-full md:block md:w-auto" id="navbar-default">
                <ul className="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white bg-gray-800 md:white:bg-gray-900 white:border-gray-700">
                    <li className="items-center align-middle flex">
                        <label className="text-xl ubuntu-bold flex items-center">{username}</label>
                    </li>
                    <li>
                        <button className="flex flex-row gap-2 items-center" onClick={()=>logoutHandler()}>
                            <LogoutIcon className="size-6" />
                            <div className="align-middle">Logout</div>
                        </button>
                    </li>
                </ul>
            </div>
        </div>
        </nav>
        </>
}