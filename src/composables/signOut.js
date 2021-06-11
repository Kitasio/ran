import axios from 'axios'
import { useRouter } from 'vue-router'

const signOut = () => {
    const router = useRouter()
    const logout = () => {
        axios
            .get("/api/logout")
            .then(response => (console.log(response))) 

        router.go()
    }

    return { logout }
}

export default signOut