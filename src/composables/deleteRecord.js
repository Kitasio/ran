import axios from 'axios'
import { useRouter } from 'vue-router'

const deleteRecord = () => {
    const router = useRouter()
    const deleteFromDb = (query, ...data) => {
        axios
            .post("http://localhost:3090/api/deleteRecord", {
                query: query,
                data: data,
            })
            .then(response => {
                router.go(-1)
            }) 
    }

    return { deleteFromDb }
}

export default deleteRecord