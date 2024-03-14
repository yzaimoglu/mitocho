import { PUBLIC_MITOCHO_API_URL } from '$env/static/public';

const baseAPIURL: string = PUBLIC_MITOCHO_API_URL;

export type BaseRes = {
  status: number,
  request_id: string,
};

type FetchType = typeof fetch;

export default class BaseService {
  private fetch: FetchType;

  constructor(fetchType: FetchType) {
    this.fetch = fetchType;
  }

  async get<O>(endpoint: string, authorization?: string): Promise<O> {
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

  async post<I, O>(endpoint: string, data: I, authorization?: string): Promise<O> {
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

  async put<I, O>(endpoint: string, data: I, authorization?: string): Promise<O> {
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

  async delete<O>(endpoint: string, authorization?: string): Promise<O> {
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
