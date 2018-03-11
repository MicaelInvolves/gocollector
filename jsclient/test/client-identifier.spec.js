/* global describe, it, beforeEach */

import chai from 'chai';
import ClientGateway from '../src/client-identifier';

chai.should();

let identifier, documentMock, dateProviderMock;

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

    identifier = new ClientGateway(documentMock, dateProviderMock);
  });

  it('should return a new uuid', () => {
    const id = identifier.id();

    id.should.be.a('string');
    id.should.have.length(36);
  });

  it('should always return the same uuid', () => {
    identifier.id().should.be.equal(identifier.id());
  });

  it('should store the id with cookies', () => {
    const id = identifier.id();

    documentMock.value.should.be.equal(`client=${id}; expires=Sat, 23 Dec 1989 14:00:00 GMT`);
  });

  it('should store the id just one time', () => {
    identifier.id();
    identifier.id();

    documentMock.setterCount.should.be.equal(1);
  });

  it('should check if id is stored already', () => {
    documentMock.value = 'client=already-stored-cookie; expires=Thu Nov 23 1989 12:00:00 GMT-0200 (-02)';

    identifier.id();
    const id = identifier.id();

    documentMock.getterCount.should.be.equal(1);
    id.should.be.equal('already-stored-cookie');
  });
});
