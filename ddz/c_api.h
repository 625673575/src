#ifndef _UTILS_CAPI_H
#define _UTILS_CAPI_H
//头文件不同目录下的文件包含的路径也不一样，包含错误的话可能会报出重定义错误
#ifdef __cplusplus
#define APP_EXTERN_C extern "C"
#else
#define APP_EXTERN_C
#endif
#include <stdint.h>

#ifdef _MSC_VER
#define APP_EXPORT __declspec(dllexport)
#define APP_C_EXPORT APP_EXTERN_C __declspec(dllexport)
#else
#define APP_EXPORT 
#define APP_C_EXPORT APP_EXTERN_C
#endif

APP_C_EXPORT uint8_t Add(uint8_t playCard2, uint8_t deskCard);
APP_C_EXPORT uint8_t AddBytes(uint8_t* playCard2, uint8_t len);
APP_C_EXPORT void DealSmoothCard(uint8_t first, uint8_t second, uint8_t * playCard0, uint8_t *playCard1, uint8_t *playCard2, uint8_t *deskCard);
APP_C_EXPORT uint8_t TrustShip(uint8_t* ret, uint8_t len, uint8_t* playCard, uint8_t lastPlayCardLen, uint8_t* lastPlayCards, uint8_t lastPlayIdentity, uint8_t curPlayIdentity);

#endif
