import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const getAuth = () => {
    const admin = ref(null)
    const router = useRouter()

    const checkAuth = () => {
        const at = localStorage.getItem('access_token')
        const rt = localStorage.getItem('refresh_token')
        const access_token = {
            headers: { Authorization: `Bearer ${at}` }
        };
        const refresh_token = {
            headers: { Authorization: `Bearer ${rt}` }
        };
        axios
            .post("http://127.0.0.1:3090/api/auth", {}, access_token)
            .then(response => {
                if (response.data) {
                    admin.value = response.data
                    console.log(admin.value)
                }
            })
            .catch(err => {
                localStorage.clear()
            })
    }
    return { admin, checkAuth }
}

export default getAuth