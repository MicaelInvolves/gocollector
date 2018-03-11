export default class AccessGateway {
  constructor(apiPath, axios, dateProvider, pathProvider) {
    this._apiPath = apiPath;
    this._axios = axios;
    this._dateProvider = dateProvider;
    this._pathProvider = pathProvider;
  }

  accessFor(clientId) {
    this._axios.post(this._apiPath, {
      clientId: clientId,
      path: this._pathProvider.get(),
      date: this._dateProvider.get()
    });
  }
}
