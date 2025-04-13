import Ajax from "@/ajax"
import { ref } from "vue"

export const useAssistance = () => {
    const assistanceList = ref<any>([])
    const getAssistanceList = async (filterValue: any) => {
        const { data } = await Ajax.post('/assistance/list', filterValue)
        assistanceList.value = data
    }
    return {
        assistanceList,
        getAssistanceList
    }
}