import axios from 'axios'
import { useRouter } from 'vue-router'
import clientPath  from './config'

const deleteRecord = (go) => {
    const router = useRouter()
    const at = localStorage.getItem('access_token')
    const access_token = {
        headers: { Authorization: `Bearer ${at}` }
    };
    const deleteFromDb = (query, ...data) => {
        axios
            .post(clientPath + "/api/deleteRecord", {
                query: query,
                data: data,
            }, access_token)
            .then(response => {
                router.go(go)
            }) 
    }

    return { deleteFromDb }
}

export default deleteRecord
