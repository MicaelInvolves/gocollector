/* global API_PATH */

import axios from 'axios';

import AccessGateway from './access-gateway';
import ClientIdentifier from './client-identifier';

class Main {
  constructor() {
    const dateProvider = {
      get() {
        return new Date();
      }
    };

    const pathProvider = {
      get() {
        return window.location.href;
      }
    };

    this._gateway = new AccessGateway(API_PATH, axios, dateProvider, pathProvider);
    this._identifier = new ClientIdentifier(document, dateProvider);
  }

  listen() {
    document.addEventListener('DOMContentLoaded', (event) => {
      this._gateway.accessFor(this._identifier.id());
    });
  }
}

new Main().listen();
