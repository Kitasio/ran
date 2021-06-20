import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import clientPath  from './config'

const updateRecord = () => {
    const router = useRouter()
    const file = ref(null)
    const fileName = ref('')
    const filePath = ref('')
    const fileUrl = ref('')
    const at = localStorage.getItem('access_token')
    const access_token = {
        headers: { Authorization: `Bearer ${at}` }
    };

    const updateToDb = async (query, ...data) => {
        console.log(query, data)
        axios
            .post(clientPath + "/api/updateRecord", {
                query: query,
                data: data,
            }, access_token)
            .then(response => {
                router.go()
            }) 
    }

    const handleChange = (e) => {
        // allowed file types
        const types = ['image/png', 'image/jpeg']

        // selected file
        const selected = e.target.files[0]

        if (selected && types.includes(selected.type)) {
            file.value = selected
            let formData = new FormData();
        
            formData.append("file", file.value);
            axios.post(clientPath + '/api/upload', formData, access_token)
            .then(res => {
                fileUrl.value = res.data["filepath"]
                console.log(res.data)
            })
            .catch(err => {
                console.log(err)
            })
        } else {
            file.value = null
        }
    }


    return { updateToDb, handleChange, fileUrl }
}

export default updateRecord
