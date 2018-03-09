/* global describe, it, beforeEach */

import chai from 'chai';
import AccessGateway from '../src/access-gateway';

chai.should();

const path = 'http://awesomeproduct.com.br/about.html';
const date = new Date(1989, 10, 23, 12, 0, 0, 0);
const apiPath = 'http://mockedapi.com';

let gateway, axiosMock, dateProviderMock, pathProvider;

describe('AccessGateway', () => {

  beforeEach(() => {
    axiosMock = {
      path: '',
      data: undefined,
      post(path, data) {
        this.path = path;
        this.data = data;
      }
    };
    dateProviderMock = {
      get() {
        return date;
      }
    };
    pathProvider = {
      get() {
        return path;
      }
    };
    gateway = new AccessGateway(apiPath, axiosMock, dateProviderMock, pathProvider);
  });

  it('should post access data', () => {
    gateway.accessFor('client-id');

    axiosMock.path.should.be.equal(apiPath);
    axiosMock.data.should.be.deep.equal({
      id: 'client-id',
      path: path,
      date: date
    });
  });
});
