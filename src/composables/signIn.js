import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import clientPath  from './config'

const signIn = () => {
    const router = useRouter()

    const username = ref('')
    const password = ref('')
    const error = ref(null)

    const login = () => {
        error.value = null
        axios
            .post(clientPath + "/api/login", {
                username: username.value,
                password: password.value, 
            })
            .then(response => {
                router.go()
                const at = response.data['access_token']
                localStorage.setItem('access_token', at)
            })
            .catch(err => {
                error.value = 'Wrong username or password'
            })
    }

    return { username, password, error, login }
}

export default signIn
