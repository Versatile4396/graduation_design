import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useTagsStore = defineStore('tags', () => {
  const tags = ref<string[]>([])
  const setTags = (_tags: string[]) => {
    tags.value = _tags
  }
  return { tags, setTags }
})

export const useTopicsStore = defineStore('topic', () => {
  const topics = ref<string[]>([])
  const setTopics = (_topic: []) => {
    topics.value = _topic
  }
  return { topics, setTopics }
})

export const useCategorieStore = defineStore('Categorie', () => {
  const Categories = ref()
  const setCategories = (_cateGorie: {}) => {
    Categories.value = _cateGorie
  }
  return { Categories, setCategories }
})
