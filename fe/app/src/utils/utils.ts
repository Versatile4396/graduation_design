// 滚动Anchor

// 点击滚动到指定位置

export const scrollToAnchor = (anchorId: string) => {
  const anchorElement = document.getElementById(anchorId)
  if (anchorElement) {
    anchorElement.scrollIntoView({ behavior: 'smooth' })
  }
}
