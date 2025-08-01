#include "reverse_string.h"

namespace reverse_string {
    std::string reverse_string(const std::string& x){
        std::string result = x;
        int mid = x.length()/2;
        for(int i = 0; i< mid;i++){
            std::swap(result[i], result[result.length() - 1 - i]);
        }
        return result;
    }
}  // namespace reverse_string
