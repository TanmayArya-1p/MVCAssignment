import * as auth from "../api/auth";
import useAuthStore from "../stores/authStore";



export default function Navbar() {
    const { role ,username} = useAuthStore.getState();



    const logoutHandler = async () => {
        if(!window.confirm("Are you sure you want to logout?")) {
            return
        }
        const {setUsername, setAuthToken, setRefreshToken,setRole} = useAuthStore.getState();

        setUsername("");
        setAuthToken("");
        setRefreshToken("");
        setRole("");

        await auth.LogoutUser();
        window.location.href = "/";
    }



    return <>
        <nav className="bg-white border-b-2 border-black">
        <div className="max-w-screen flex flex-wrap items-center justify-between mx-auto p-4">
            <a href="/home" className="flex items-center space-x-3 rtl:space-x-reverse">
                <span className="self-center text-3xl ubuntu-bold">InOrder</span>
                <span className="flex items-end text-lg ubuntu-regular self-end">{role == "customer" ? "" : role}</span>
            </a>
            <div className="flex flex-row gap-4 text-lg">
                    {role == "admin" && <>
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
                <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
                    <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M1 1h15M1 7h15M1 13h15"/>
                </svg>
            </button>
            <div className="hidden w-full md:block md:w-auto" id="navbar-default">
                <ul className="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white bg-gray-800 md:white:bg-gray-900 white:border-gray-700">
                    <li className="items-center align-middle flex">
                        <label className="text-xl ubuntu-bold flex items-center">{username}</label>
                    </li>
                    <li>
                        <button className="flex flex-row gap-2 items-center" onClick={()=>logoutHandler()}>
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" style={{ width: '1.2rem', height: '1.2rem', display: 'block' }} className="align-middle"><path d="M160 96c17.7 0 32-14.3 32-32s-14.3-32-32-32L96 32C43 32 0 75 0 128L0 384c0 53 43 96 96 96l64 0c17.7 0 32-14.3 32-32s-14.3-32-32-32l-64 0c-17.7 0-32-14.3-32-32l0-256c0-17.7 14.3-32 32-32l64 0zM502.6 278.6c12.5-12.5 12.5-32.8 0-45.3l-128-128c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L402.7 224 192 224c-17.7 0-32 14.3-32 32s14.3 32 32 32l210.7 0-73.4 73.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0l128-128z"/></svg>
                            <div className="align-middle">Logout</div>
                        </button>
                    </li>
                </ul>
            </div>
        </div>
        </nav>
        </>
}