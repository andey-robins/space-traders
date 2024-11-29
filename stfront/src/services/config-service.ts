import type { Config } from "@/types/config";
import axios, { type AxiosResponse } from "axios";

class ConfigService {
    private baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/configs"
    }

    async getAll(): Promise<AxiosResponse<Config[]>> {
        console.log('invoked get all')
        return axios.get<Config[]>(`${this.baseUrl}/all`)
    }

    async save(dto: Config): Promise<void> {
        return axios.post(`${this.baseUrl}/save`, dto)
    }
}

export const configService = new ConfigService;