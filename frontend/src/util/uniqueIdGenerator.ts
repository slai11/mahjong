
let i = Math.floor(Math.random() * 4e9)

export default function uniqueIdGenerator () {
  i++
  return i.toString().padStart(9, '0')
}
