import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import clientPath  from './config'

const getAuth = () => {
    const isAdmin = ref(null)
    const username = ref('')
    const uid = ref(null)
    const router = useRouter()

    const checkAuth = () => {
        const at = localStorage.getItem('access_token')
        const access_token = {
            headers: { Authorization: `Bearer ${at}` }
        };
        axios
            .post(clientPath + "/api/auth", {}, access_token)
            .then(response => {
                if (response.data) {
                    isAdmin.value = response.data.admin
                    username.value = response.data.username
                    uid.value = response.data.uid
                }
            })
            .catch(err => {
                localStorage.clear()
            })
    }
    return { isAdmin, username, uid, checkAuth }
}

export default getAuth
