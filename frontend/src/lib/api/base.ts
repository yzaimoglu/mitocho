const baseAPIURL: string = 'http://localhost:3000/api/v1/';

type FetchType = typeof fetch;

interface IServiceResponse<T = any> {
  [key: string]: T;
}

export default class BaseService {
  private fetch: FetchType;

  constructor(fetchType: FetchType) {
    this.fetch = fetchType;
  }

  async get<T>(endpoint: string, authorization?: string): Promise<IServiceResponse<T>> {
    let response: any;
    if (typeof authorization !== 'undefined') {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${authorization}`
        },
      });
    } else {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'GET',
        headers: {
          Accept: 'application/json',
        },
      });
    }
    return response.json();
  }

  async post<T>(endpoint: string, data: T, authorization?: string): Promise<IServiceResponse> {
    let response: any;
    if (typeof authorization !== 'undefined') {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${authorization}`
        },
        body: JSON.stringify(data),
      });
      
    } else {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
    }
    return response.json();
  }

  async put<T>(endpoint: string, data: T, authorization?: string): Promise<IServiceResponse> {
    let response: any;
    if (typeof authorization !== 'undefined') {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${authorization}`
        },
        body: JSON.stringify(data),
      });
      
    } else {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
    }
    return response.json();
  }

  async delete(endpoint: string, authorization?: string): Promise<IServiceResponse> {
    let response: any;
    if (typeof authorization !== 'undefined') {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${authorization}`
        },
      });
      
    } else {
      response = await this.fetch(`${baseAPIURL}${endpoint}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });
    }
    return response.json();
  }
}
