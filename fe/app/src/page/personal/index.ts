import Ajax from "@/ajax"
import type { UserInfo } from "@/ajax/type/user"
import { ref } from "vue"

export const useAssistance = () => {
    const assistanceList = ref<any>([])
    const getAssistanceList = async (filterValue: any = {}) => {
        const { data } = await Ajax.post('/assistance/list', filterValue)
        assistanceList.value = data
    }
    return {
        assistanceList,
        getAssistanceList
    }
}

export const useFollower = () => {
    const followerList = ref<UserInfo[]>()
    const followedList = ref<UserInfo[]>()
    const getFollowerList = async (uid?: number) => {
        const { data } = await Ajax.post('/follow/list/' + uid,)
        followerList.value = data.followList || []
        followedList.value = data.followedList || []
    }
    return {
        followerList,
        followedList,
        getFollowerList,
    }
}

