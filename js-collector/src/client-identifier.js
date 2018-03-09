import uuid from 'uuid/v4';

function getCookie(cname, document) {
  let name = `${cname}=`;
  let cookies = document.cookie.split(';');

  for (let i = 0; i < cookies.length; i++) {
    let cookieValue = cookies[i];

    while (cookieValue.charAt(0) === ' ') {
      cookieValue = cookieValue.substring(1);
    }
    if (cookieValue.indexOf(name) === 0) {
      return cookieValue.substring(name.length, cookieValue.length);
    }
  }

  return '';
}

export default class ClientGateway {

  constructor(document, dateProvider) {
    this._document = document;
    this._dateProvider = dateProvider;
  }

  id() {
    if (!this._id) {
      let storedId = getCookie('client', this._document);

      this._id = storedId || uuid();
      this._document.cookie = `client=${this._id}; expires=${this._dateProvider.get()}`;
    }
    return this._id;
  }

}
