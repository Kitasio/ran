import { ref } from 'vue'
import axios from 'axios'

const getAuth = () => {
    const admin = ref(null)

    const checkAuth = () => {
        try {
            axios
                .get("/api/auth")
                .then(response => {
                    if (response.data) {
                        admin.value = response.data
                    }
                })
        } catch(err) {
            console.log(err)
        }
    }
    return { admin, checkAuth }
}

export default getAuth