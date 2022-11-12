class LocalObjectStorage {
  constructor() {
    if (!window.localStorage) {
      throw new Error('Unsupported local storage')
    }
  }

  set(key, data) {
    localStorage.setItem(key, JSON.stringify(data))
  }

  get(key) {
    return JSON.parse(localStorage.getItem(key))
  }
}

const storage = new LocalObjectStorage();

export default storage
