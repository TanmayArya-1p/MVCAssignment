import Spinner from "@/components/spinner";
import useAuthStore from "@/stores/authStore";
import { roles } from "@/utils/const";

export default function AdminProtectedRoute({children}) {
    const {role} = useAuthStore.getState();

    if(role !== roles.ADMIN) {
        window.location.pathname = "/notfound";
        return <div className="h-screen w-screen flex items-center justify-center">
            <Spinner />
        </div>
    }

    return <>
        {children}
    </>
}