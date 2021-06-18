import axios from 'axios'
import { useRouter } from 'vue-router'

const signOut = () => {
    const router = useRouter()
    const logout = () => {
        localStorage.clear()
        router.go()
    }

    return { logout }
}

export default signOut