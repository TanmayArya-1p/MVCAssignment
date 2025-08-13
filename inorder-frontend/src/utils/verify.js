import * as auth from '@/api/auth';
import useAuthStore from '@/stores/authStore';
import { ONBOARDING_PATHS } from './const';


export default async function VerifySignedIn() {
    try {
        let resp =await auth.RefreshToken();
        const {setAuthToken, setRefreshToken} = useAuthStore.getState();
        setAuthToken(resp.authToken);
        setRefreshToken(resp.refreshToken);
    } catch (error) {
        if(!ONBOARDING_PATHS.includes(window.location.pathname)) {
            window.location.pathname = "/";
        }
        throw error
    }
}