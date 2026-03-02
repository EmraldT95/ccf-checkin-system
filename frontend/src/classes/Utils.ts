export class Utils {
  static debounce(callback: CallableFunction, duration = 700) {
    let timer = 0

    return (...args: any) => {
      clearTimeout(timer)
      timer = setTimeout(() => callback(...args), duration)
    }
  }

  static generateRandomString(length: number): string {
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    let result = ''

    for (let i = 0; i < length; i++) {
      const randomIndex = Math.floor(Math.random() * characters.length)
      result += characters.charAt(randomIndex)
    }

    return result
  }
}
