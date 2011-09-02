package errforce

import "os"

func IsEPERM(err os.Error) bool { return Is(os.EPERM, err) }
func IsENOENT(err os.Error) bool { return Is(os.ENOENT, err) }
func IsESRCH(err os.Error) bool { return Is(os.ESRCH, err) }
func IsEINTR(err os.Error) bool { return Is(os.EINTR, err) }
func IsEIO(err os.Error) bool { return Is(os.EIO, err) }
func IsENXIO(err os.Error) bool { return Is(os.ENXIO, err) }
func IsE2BIG(err os.Error) bool { return Is(os.E2BIG, err) }
func IsENOEXEC(err os.Error) bool { return Is(os.ENOEXEC, err) }
func IsEBADF(err os.Error) bool { return Is(os.EBADF, err) }
func IsECHILD(err os.Error) bool { return Is(os.ECHILD, err) }
func IsEDEADLK(err os.Error) bool { return Is(os.EDEADLK, err) }
func IsENOMEM(err os.Error) bool { return Is(os.ENOMEM, err) }
func IsEACCES(err os.Error) bool { return Is(os.EACCES, err) }
func IsEFAULT(err os.Error) bool { return Is(os.EFAULT, err) }
func IsEBUSY(err os.Error) bool { return Is(os.EBUSY, err) }
func IsEEXIST(err os.Error) bool { return Is(os.EEXIST, err) }
func IsEXDEV(err os.Error) bool { return Is(os.EXDEV, err) }
func IsENODEV(err os.Error) bool { return Is(os.ENODEV, err) }
func IsENOTDIR(err os.Error) bool { return Is(os.ENOTDIR, err) }
func IsEISDIR(err os.Error) bool { return Is(os.EISDIR, err) }
func IsEINVAL(err os.Error) bool { return Is(os.EINVAL, err) }
func IsENFILE(err os.Error) bool { return Is(os.ENFILE, err) }
func IsEMFILE(err os.Error) bool { return Is(os.EMFILE, err) }
func IsENOTTY(err os.Error) bool { return Is(os.ENOTTY, err) }
func IsEFBIG(err os.Error) bool { return Is(os.EFBIG, err) }
func IsENOSPC(err os.Error) bool { return Is(os.ENOSPC, err) }
func IsESPIPE(err os.Error) bool { return Is(os.ESPIPE, err) }
func IsEROFS(err os.Error) bool { return Is(os.EROFS, err) }
func IsEMLINK(err os.Error) bool { return Is(os.EMLINK, err) }
func IsEPIPE(err os.Error) bool { return Is(os.EPIPE, err) }
func IsEAGAIN(err os.Error) bool { return Is(os.EAGAIN, err) }
func IsEDOM(err os.Error) bool { return Is(os.EDOM, err) }
func IsERANGE(err os.Error) bool { return Is(os.ERANGE, err) }
func IsEADDRINUSE(err os.Error) bool { return Is(os.EADDRINUSE, err) }
func IsECONNREFUSED(err os.Error) bool { return Is(os.ECONNREFUSED, err) }
func IsENAMETOOLONG(err os.Error) bool { return Is(os.ENAMETOOLONG, err) }
func IsEAFNOSUPPORT(err os.Error) bool { return Is(os.EAFNOSUPPORT, err) }
func IsETIMEDOUT(err os.Error) bool { return Is(os.ETIMEDOUT, err) }
func IsENOTCONN(err os.Error) bool { return Is(os.ENOTCONN, err) }
