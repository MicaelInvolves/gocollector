/* global describe, it, beforeEach */

import chai from 'chai';
import ClientGateway from '../src/client-identifier';

chai.should();

let gateway, documentMock, dateProviderMock;

describe('ClientGateway', () => {

  beforeEach(() => {
    documentMock = {
      value: '',
      setterCount: 0,
      getterCount: 0,
      set cookie(value) {
        this.setterCount += 1;
        this.value = value;
      },
      get cookie() {
        this.getterCount += 1;
        return this.value;
      }
    };

    dateProviderMock = {
      get() {
        return new Date(1989, 10, 23, 12, 0, 0, 0);
      }
    };

    gateway = new ClientGateway(documentMock, dateProviderMock);
  });

  it('should return a new uuid', () => {
    const id = gateway.id();

    id.should.be.a('string');
    id.should.have.length(36);
  });

  it('should always return the same uuid', () => {
    gateway.id().should.be.equal(gateway.id());
  });

  it('should store the id with cookies', () => {
    const id = gateway.id();

    documentMock.value.should.be.equal(`client=${id}; expires=Thu Nov 23 1989 12:00:00 GMT-0200 (-02)`);
  });

  it('should store the id just one time', () => {
    gateway.id();
    gateway.id();

    documentMock.setterCount.should.be.equal(1);
  });

  it('should check if id is stored already', () => {
    documentMock.value = 'client=already-stored-cookie; expires=Thu Nov 23 1989 12:00:00 GMT-0200 (-02)';

    gateway.id();
    const id = gateway.id();

    documentMock.getterCount.should.be.equal(1);
    id.should.be.equal('already-stored-cookie');
  });
});
