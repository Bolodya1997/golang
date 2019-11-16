#ifndef SEARKING_GOLANG_GO_OS_SIGNAL_CGO_SIGNAL_CGO_H_
#define SEARKING_GOLANG_GO_OS_SIGNAL_CGO_SIGNAL_CGO_H_
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif
void PrintStr(char *s);

int CGOSignalHandlerSignalAction(int signum);
void CGOSignalHandlerSetFd(int fd);
void CGOSignalHandlerSetBacktraceDump(bool enable);

#ifdef __cplusplus
}
#endif

#endif // SEARKING_GOLANG_GO_OS_SIGNAL_CGO_SIGNAL_CGO_H_