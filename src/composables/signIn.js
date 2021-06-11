import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const signIn = () => {
    const router = useRouter()

    const username = ref('')
    const password = ref('')
    const error = ref(null)

    const login = () => {
        error.value = null
        axios
            .post("/api/login", {
                name: username.value,
                password: password.value,
            })
            .then(response => {
                router.go()
            })
            .catch(err => {
                error.value = 'Wrong username or password'
            })
    }

    return { username, password, error, login }
}

export default signIn